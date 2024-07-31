import Home from "./pages/Home";
import TasksList from "./pages/TasksList";
import Header from "./components/Header";

function App() {

  return (
    <div className="container">

      {/* Header */}
      <div className="row">
        <Header />
      </div >

      {/* Body */}
      < div className="row" >
        <Home />
        <TasksList />



      </div >
    </div >
  );
}

export default App;
