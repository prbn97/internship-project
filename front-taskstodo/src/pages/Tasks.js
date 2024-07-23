import { useEffect, useState } from "react";
import { Link, useOutletContext } from "react-router-dom";

import logo from '../img/logo.svg';
import Title from '../components/Title';
import TasksList from "../components/TasksList";


const Tasks = () => {
	const { user } = useOutletContext(); // get user context 

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
