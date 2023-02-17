import axios from "axios";
import { serverURL } from "src/config/constants";
import { ICertification } from "src/model/interfaces";

export interface IUseGetFromCredlyLogic{
    validate:(values:any)=>any;
    initialValues: any;
    onSubmit: (values:any, {setSubmitting}:any)=>void
}   

function UseGetFromCredlyLogic(props:any):IUseGetFromCredlyLogic{
    const initialValues = { url:'' }
    const api_url = `${serverURL}/api/credly`;
    
    function validate(values:any):any{
        const errors:any = {};
        if (!values.url) {
            errors.url = 'Required';
        } 
        return errors;
    }

    async function onSubmit(values:any, { setSubmitting }:any):Promise<void> {

        setSubmitting(true)
        const data = JSON.stringify(values)
        const response = await axios.post(api_url, data, {
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
export default UseGetFromCredlyLogic;