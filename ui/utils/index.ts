export function capitalize(value: string): string {
    return value.charAt(0).toUpperCase() + value.slice(1);
}

export function weiToGwei(valueInWei: string): string {
    const gwei: number = Number(valueInWei) / 1000000000;
    return gwei.toFixed(2);
}

export function timeDiffInSeconds(time: Date): number {
    const now: Date = new Date();
    return Math.round((now.getTime() - time.getTime()) / 1000);
}


export function noop(): void {
    // do nothing
}