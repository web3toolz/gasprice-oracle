import {useEffect, useState} from "react";
import {useInterval} from '@mantine/hooks';
import {timeDiffInSeconds} from "@/utils";

interface Props {
    updateTime?: Date;
}

export default function Counter({updateTime}: Props) {
    const [timeDiff, setTimeDiff] = useState<number>(0);

    // update counter every second
    const timerInternal = useInterval(() => setTimeDiff(timeDiffInSeconds(updateTime || null)), 1000);

    useEffect(() => {
        timerInternal.start();
        return timerInternal.stop;
    }, [updateTime, timerInternal])

    return (
        <div className="text-white">Last update: {timeDiff} seconds ago</div>
    );
}
