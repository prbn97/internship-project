const Title = ({ icon, text }) => {
    return (
        <div className="d-flex mt-2">
            <img
                alt="icon-title"
                src={icon}
                width="30"
                height="30"
                className="me-1"
            />
            <h4>
                {text}
            </h4>
        </div>
    );
}

export default Title;