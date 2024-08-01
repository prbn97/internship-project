
import { useState } from "react";
import { Outlet } from "react-router-dom";

import Header from "./components/Header";
import Alert from "./components/Alert"

function App() {
  const [jwtToken, setJwtToken] = useState("");

  // state variables to store infos to alert component
  const [alertMessage, setAlertMessage] = useState("");
  const [alertClassName, setAlertClassName] = useState("d-none"); //d-none bootstrap class to display:none



  return (
    <div className="container">

      <div className="row">
        <Header jwtToken={jwtToken} setJwtToken={setJwtToken} />
      </div >

      < div className="row" >
        {/*inside  <Outlet/>  we display the components, with "context" you can pass informatin like jwt token and alert*/}
        <Alert message={alertMessage} className={alertClassName} />
        <Outlet context={{
          jwtToken,
          setJwtToken,
          setAlertClassName,
          setAlertMessage
        }} />


      </div >
    </div >
  );
}

export default App;
