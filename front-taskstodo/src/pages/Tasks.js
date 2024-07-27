

import logo from '../img/logo.svg';
import Title from '../components/Title';
import TasksList from "../components/TasksList";


const Tasks = () => {

	return (
		<>
			<div className="row ">
				<Title icon={logo} text="List of Tasks" />
			</div>


			<div className="row ">
				<TasksList />
			</div>
		</>
	);
};

export default Tasks;
