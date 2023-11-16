export function capitalize(value: string): string {
    return value.charAt(0).toUpperCase() + value.slice(1);
}

export function weiToGwei(valueInWei: string): string {
    const gwei: number = Number(valueInWei) / 1000000000;
    return gwei.toFixed(2);
}

export function timeDiffInSeconds(time: Date | null): number {
    if (!time) {
        return -1;
    }
    const now: Date = new Date();
    return Math.round((now.getTime() - time.getTime()) / 1000);
}

// ethereum-mainnet to Ethereum Mainnet
export function formatNetworkName(networkName: string): string {
    return networkName.replace(/-/g, ' ').replace(/\b\w/g, (l) => l.toUpperCase());
}

export function noop(): void {
    // do nothing
}