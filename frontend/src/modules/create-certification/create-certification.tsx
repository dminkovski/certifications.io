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
import UseGetFromCredlyLogic from './get-from-credly.logic';

function CreateCertification(props?:any) {

  const createCertification = UseCreateCertificationLogic(props);
  const getFromCredly = UseGetFromCredlyLogic(props);
 
  return (
    <div id="create-certification">
      <h1>
        Get from Credly
      </h1>
     <Formik
       initialValues={getFromCredly.initialValues}
       validate={getFromCredly.validate}
       onSubmit={getFromCredly.onSubmit}
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
           <TextInput
             onChange={handleChange}
             onBlur={handleBlur}
             value={values.url}
             label="Credly url" name="url" description="URL of the certification on credly" />
            {errors.url && touched.url && errors.url}
           <button type="submit" disabled={isSubmitting}>
             Submit
           </button>
           </>
         </form>
       )}
     </Formik>
     <h1>Create New Certification</h1>
     <Formik
       initialValues={createCertification.initialValues}
       validate={createCertification.validate}
       onSubmit={createCertification.onSubmit}
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
           <TextInput
             onChange={handleChange}
             onBlur={handleBlur}
             value={values.name}
             label="Certification name" name="name" description="Official name of the certification" />
            {errors.name && touched.name && errors.name}
           <TextInput
             onChange={handleChange}
             onBlur={handleBlur}
             value={values.link}
             label="Certification link" name="link" description="Official link of the certification on credly" />
            {errors.link && touched.link && errors.link}
           <TextInput
             onChange={handleChange}
             onBlur={handleBlur}
             value={values.image}
             label="Certification image" name="image" description="Image URL on credly" />
            {errors.image && touched.image && errors.image}
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
