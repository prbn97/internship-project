
import Header from "./components/Header";
import { Outlet } from "react-router-dom";

function App() {

  return (
    <div className="container">

      {/* Header */}
      <div className="row">
        <Header />
      </div >

      {/* Body */}
      < div className="row" >

        <Outlet />
        {/*inside  <Outlet/>  we display the components*/}


      </div >
    </div >
  );
}

export default App;
