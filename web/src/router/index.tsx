import {createBrowserRouter} from "react-router-dom";
import App from "@/App.tsx";
import AuthPage from "@/features/auth";
import AuthRouter from "@/router/AuthRouter.tsx";

export const router = createBrowserRouter([
    {
        path:"/auth",
        element:<AuthPage/>
    },
    {
        element:<AuthRouter/>,
        children:[
            {
                index:true,
                element:<App/>
            }
        ]
    }
])