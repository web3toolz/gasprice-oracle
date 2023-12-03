import {capitalize, weiToGwei} from "@/utils";

interface Props {
    title: string;
    subtitle?: string;
    gasPriceValue: string;
}


export default function Card({title, subtitle, gasPriceValue}: Props) {

    let amount, unit;

    if (parseInt(gasPriceValue) > 10 ** 9) {
        amount = weiToGwei(gasPriceValue);
        unit = 'gwei';
    } else {
        amount = gasPriceValue;
        unit = 'wei';
    }

    return (
        <div className="rounded-lg shadow-sm divide-y divide-zinc-600 bg-zinc-900">
            <div className="p-6">
                <h2 className="text-2xl font-semibold leading-6 text-white">{capitalize(title)}</h2>
                {subtitle && <p className="mt-4 text-zinc-300">Subtext</p>}
                <p className="mt-8">
                    {
                        gasPriceValue ?
                            <span className="text-5xl font-extrabold text-white">{amount}</span> :
                            <span className="text-5xl font-extrabold text-white">?</span>
                    }
                    <span className="text-base font-medium text-zinc-100 ml-2">{unit}</span>
                </p>
            </div>
        </div>
    )
}