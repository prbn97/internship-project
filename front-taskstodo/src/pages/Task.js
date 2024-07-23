import { useEffect, useState } from "react";
import { useParams, useOutletContext, useNavigate } from "react-router-dom";
import logo from '../img/logo.svg';
import edit from '../img/edit.svg';
import close from '../img/close.svg';
import save from '../img/save.svg';
import del from '../img/delete.svg';
import Title from '../components/Title';
import TaskEdition from "../components/TaskEdition";

const Task = () => {
    const { user } = useOutletContext(); // get user mock login context 


    return (
        <>
            {/* Bread-crumb */}
            <p>tasks-task</p>

            <TaskEdition />
        </>
    );
};

export default Task;
