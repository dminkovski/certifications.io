package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/dminkovski/certifications.io/api/model"
)

const colKey = "Certifications"

// Insert Certifications into Database
func InsertCertifications(certifications []model.Certification) int {
	db, err := Connect()
	if err != nil {
		log.Panic(err)
	}
	col := db.GetDatabase().Collection(colKey)

	data := make([]interface{}, 0)
	for i, _ := range certifications {
		data = append(data, interface{}(
			certifications[i],
		))
	}
	result, err := col.InsertMany(context.TODO(), data)
	if err != nil {
		log.Panic(err)
	}
	for _, id := range result.InsertedIDs {
		fmt.Printf("Inserted Certification with _id: %v\n", id)
	}
	err = db.Disconnect()
	if err != nil {
		log.Panic(err)
	}
	return len(result.InsertedIDs)
}

// Retrieve Certifications from DB
func GetCertifications() []model.Certification {
	db, err := Connect()
	if err != nil {
		log.Panic(err)
	}
	col := db.GetDatabase().Collection(colKey)
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Panic(err)
	}
	var results []model.Certification
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	if err := cursor.Err(); err != nil {
		log.Panic(err)
	}
	err = db.Disconnect()
	if err != nil {
		log.Panic(err)
	}
	return results
}

// Saves the Course to the Certification ID
func SaveCourse(course model.Course, certId string) error {
	db, err := Connect()
	if err != nil {
		log.Panic(err)
	}
	col := db.GetDatabase().Collection(colKey)
	objectId, err := primitive.ObjectIDFromHex(certId)
	if err != nil{
		log.Println("Invalid id")
	}
	result := col.FindOne(context.TODO(), bson.M{
		"_id": objectId,
	})
	if err != nil {
		log.Panic(err)
		return err
	}
	certification := model.Certification{}
	result.Decode(&certification)
	if certification.AddCourse(course) {
		filter := bson.D{{"_id", objectId}}
		update := bson.D{{"$set", bson.M{"courses":certification.Courses, "updated": certification.Updated}}}

		col.UpdateOne(
			context.TODO(),
			filter,
			update,
		)
	}
	
	err = db.Disconnect()
	if err != nil {
		log.Panic(err)
		return err
	}	
	return nil
}