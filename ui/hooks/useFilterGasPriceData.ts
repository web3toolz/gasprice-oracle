import {NetworkData} from "@/types";
import {useEffect, useState} from "react";


const defaultChosenNetwork: string = "ethereum-mainnet";


export function useFilterGasPriceData(networksData: NetworkData[]) {
    const [chosenNetwork, setChosenNetwork] = useState<string | undefined>(defaultChosenNetwork);
    const [networkData, setNetworkData] = useState<NetworkData | undefined>();

    useEffect(() => {
        if (chosenNetwork) {
            setNetworkData(networksData.filter((networkData: NetworkData) => networkData.title === chosenNetwork)[0]);
        }
    }, [chosenNetwork, networksData]);

    return {chosenNetwork, setChosenNetwork, networkData};

}