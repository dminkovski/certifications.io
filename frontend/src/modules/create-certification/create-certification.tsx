import { IconCheck, IconExternalLink, IconStar } from "@tabler/icons";

import {
  AppShell,
  Card,
  Navbar,
  Container,
  Header,
  NavLink,
  Text,
  Title,
  Badge,
  Image,
  Group,
  Table,
  Button,
  TextInput,
} from "@mantine/core";

import { Formik } from 'formik';
import UseCreateCertificationLogic  from './create-certification.logic';

function CreateCertification(props?:any) {

  const {validate, initialValues, onSubmit} = UseCreateCertificationLogic(props);
 
  return (
    <div id="create-certification">
     <h1>Create Certification</h1>
     <Formik
       initialValues={initialValues}
       validate={validate}
       onSubmit={onSubmit}
     >
       {({
         values,
         errors,
         touched,
         handleChange,
         handleBlur,
         handleSubmit,
         isSubmitting,
       }) => (
         <form onSubmit={handleSubmit}>
           <>
           <TextInput              onChange={handleChange}
             onBlur={handleBlur}
             value={values.name}
             label="Certification name" name="name" description="Official name of the certification" />
           {errors.name && touched.name && errors.name}
           <button type="submit" disabled={isSubmitting}>
             Submit
           </button>
           </>
         </form>
       )}
     </Formik>
    </div>
  );
}
export default CreateCertification;
