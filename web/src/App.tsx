import { Button } from "@/components/ui/button";
import { useGetHealth } from "@/api/default";

export default function App() {
    const { data, isPending, refetch } = useGetHealth();

    if (isPending) {
        return <div>正在检查服务状态...</div>;
    }

    if (!data) {
        return <div>暂无数据</div>;
    }

    if (data.status !== 200) {
        return <div>服务异常：{data.data.detail}</div>;
    }

    return (
        <div>
            <p>状态码：{data.status}</p>
            <p>服务状态：{data.data.message}</p>

            <Button onClick={() => void refetch()}>
                重新检查
            </Button>
        </div>
    );
}