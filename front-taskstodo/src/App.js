import { useState } from 'react';
import { Outlet } from "react-router-dom";
import Header from "./components/Header";

function App() {
  const [jwtToken, setJwtToken] = useState("");
  const [user, setUser] = useState(null);
  const [tasks, setTasks] = useState([]);

  return (
    <div className="Container">
      <header className="row">
        <Header jwtToken={jwtToken} setJwtToken={setJwtToken} user={user} setUser={setUser} tasks={tasks} setTasks={setTasks} />
      </header>
      <main className="row justify-content-center">
        <div className="col-6">
          <Outlet context={{ user, setUser, setJwtToken, tasks, setTasks }} />
        </div>
      </main>
    </div>
  );
}

export default App;
