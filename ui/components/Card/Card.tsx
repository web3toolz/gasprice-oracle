import {capitalize, weiToGwei} from "@/utils";

interface Props {
    title: string;
    gasPriceValue: string;
}


export default function Card({title, gasPriceValue}: Props) {

    return (
        <div className="rounded-lg shadow-sm divide-y divide-zinc-600 bg-zinc-900">
            <div className="p-6">
                <h2 className="text-2xl font-semibold leading-6 text-white">{capitalize(title)}</h2>
                <p className="mt-4 text-zinc-300">Subtext</p>
                <p className="mt-8">
                    <span className="text-5xl font-extrabold white">{weiToGwei(gasPriceValue)}</span>
                    <span className="text-base font-medium text-zinc-100 ml-2">gwei</span>
                </p>
            </div>
        </div>
    )
}