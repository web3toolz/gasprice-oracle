import {ComboboxItem, NativeSelect} from '@mantine/core';
import {formatNetworkName} from "@/utils";

interface Props {
    value: string;
    options: string[];
    onChange: Function;
}

export default function NetworkSelector({value, options, onChange}: Props) {
    const formattedOptions: ComboboxItem[] = options.map((item: string) => {
        return {
            label: formatNetworkName(item),
            value: item,
        }
    })

    return <NativeSelect
        classNames={{
            description: "",
            input: "",
        }}
        value={value}
        label="Select network"
        data={formattedOptions}
        onChange={(event) => onChange(event.currentTarget.value)}
    />;
}