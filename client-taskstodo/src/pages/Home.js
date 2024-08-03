import { useOutletContext } from "react-router-dom";

import LoginReminder from "../components/LoginReminder";
import TasksList from "./TasksList";
const Home = () => {

    const { jwtToken } = useOutletContext();

    return (
        <>
            {jwtToken === ""
                ? < div className="text-center"><LoginReminder /> </div >
                : <TasksList />
            }
        </>
    );
}

export default Home;