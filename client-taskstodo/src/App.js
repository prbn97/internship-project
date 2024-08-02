import { useState } from "react";
import { Outlet } from "react-router-dom";
import Header from "./components/Header";
import Alert from "./components/Alert";

function App() {
  const [jwtToken, setJwtToken] = useState("");
  const [alertMessage, setAlertMessage] = useState("");
  const [alertClassName, setAlertClassName] = useState("d-none");
  const [tasksUpdate, setTasksUpdate] = useState(0);

  const handleTaskCreated = () => {
    setTasksUpdate(prev => prev + 1);
  };

  return (
    <div className="container">
      <div className="row">
        <Header jwtToken={jwtToken} setJwtToken={setJwtToken} onTaskCreated={handleTaskCreated} />
      </div>
      <div className="row">
        <Alert message={alertMessage} className={alertClassName} />
        <Outlet context={{
          jwtToken,
          setJwtToken,
          setAlertClassName,
          setAlertMessage,
          tasksUpdate
        }} />
      </div>
    </div>
  );
}

export default App;
