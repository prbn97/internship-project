import LoginReminder from "../components/LoginReminder";
import TasksList from "./TasksList";
const Home = () => {
    return (

        <>
            <div className="text-center">
                <LoginReminder />
            </div>

            <TasksList />
        </>
    );
}

export default Home;