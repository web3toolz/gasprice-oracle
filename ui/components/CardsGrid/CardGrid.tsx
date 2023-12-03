import Card from "@/components/CardsGrid/Card";
import {ReactElement} from "react";
import {NetworkData} from "@/types";

interface Props {
    networkData?: NetworkData;
}


export default function CardGrid({networkData}: Props) {

    const cards: ReactElement[] | "" = (
        networkData && networkData.data ?
            networkData.data.map((item) => {
                    return <Card title={item.title} gasPriceValue={item.value} key={item.title}/>
                }
            )
            : ""
    )

    return (
        <div
            className="mt-12 space-y-4 sm:mt-16 sm:space-y-0 sm:grid sm:grid-cols-2 sm:gap-6 lg:max-w -4xl lg:mx-auto xl:max-w-none xl:mx-0 xl:grid-cols-3">
            {cards}
        </div>
    )

}