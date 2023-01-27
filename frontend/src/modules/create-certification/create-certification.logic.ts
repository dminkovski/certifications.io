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
    courses:  [] }
    

    function validate(values:any):any{
        const errors:any = {};
        if (!values.email) {
            errors.name = 'Required';
        } 
        return errors;
    }

    function onSubmit(values:any, { setSubmitting }:any):void{
        setTimeout(() => {
          alert(JSON.stringify(values, null, 2));
          setSubmitting(false);
        }, 400);
    }


    return {
        validate,
        initialValues,
        onSubmit
    }

}
export default UseCreateCertificationLogic;