import { useRouteError } from "react-router-dom";
import Header from "../components/Header";

export default function ErrorPage() {
    const error = useRouteError();

    return (
        <div className="container">
            {/* Header */}
            <div className="row">
                <Header />
            </div >
            <div className="row">
                <div className="col-md-6 offset-md3">
                    <h1 className="mt-3">Oops!</h1>
                    <p>Sorry, an unexpected error has occurred. </p>
                    <p>
                        <em>{error.statusText || error.message}</em>
                    </p>

                </div>
            </div>
        </div>
    )
}