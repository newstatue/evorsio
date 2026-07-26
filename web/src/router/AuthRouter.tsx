import {useAuthStore} from "@/store/useAuthStore.ts";
import {Navigate, Outlet, useLocation} from "react-router-dom";

export default function AuthRouter(){
    const token = useAuthStore((state)=>state.token)
    const location = useLocation()

    if(!token){
        return(
            <Navigate to={'/auth'}
                replace={true}
                state={{ from: location }}
            />
        )
    }

    return<Outlet/>
}