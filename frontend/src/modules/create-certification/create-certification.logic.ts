import axios from "axios";
import { serverURL } from "src/config/constants";
import { ICertification } from "src/model/interfaces";

export interface IUseCreateCertificationLogic{
    validate:(values:any)=>any;
    initialValues: any;
    onSubmit: (values:any, {setSubmitting}:any)=>void
}   

function UseCreateCertificationLogic(props:any):IUseCreateCertificationLogic{
    const initialValues:ICertification = { id: -1 ,name: '', notes: '',
    image: '',
    skills: [],
    link: '',
    minMonths:0,
    maxMonths:1,
    shortName: '',
    updated:'',
    created:'',
    price: 0,
    courses:  [] }
    const url = `${serverURL}/api/certification`;
    
    function validate(values:any):any{
        const errors:any = {};
        if (!values.name) {
            errors.name = 'Required';
        } 
        return errors;
    }

    async function onSubmit(values:any, { setSubmitting }:any):Promise<void> {

        setSubmitting(true)
        const data = JSON.stringify(values)
        const response = await axios.post(url, data, {
           headers:{ "Content-Type":"application/json"},
        });
        console.log(response)
        setSubmitting(false)
    }


    return {
        validate,
        initialValues,
        onSubmit
    }

}
export default UseCreateCertificationLogic;