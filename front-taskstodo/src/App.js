import { useState } from 'react';
import { Outlet } from "react-router-dom";
import Header from "./components/Header";

function App() {
  const [jwtToken, setJwtToken] = useState("");
  const [user, setUser] = useState(null);

  return (
    <div className="Container">
      <div className="row">
        <Header jwtToken={jwtToken} setJwtToken={setJwtToken} user={user} setUser={setUser} />
      </div>
      <div className="row justify-content-center">
        <div className="col-6">
          <Outlet context={{ user, setUser, setJwtToken }} />
        </div>
      </div>
    </div>
  );
}

export default App;
