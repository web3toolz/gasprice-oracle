export const GasPriceStrategy = {
    SLOW: "slow",
    NORMAL: "normal",
    FAST: "fast",
    FASTEST: "fastest",
};

export interface GasPriceData {
    title: string;
    value: string;
}

export interface NetworkData {
    title: string;
    updatedAt?: Date;
    data?: GasPriceData[];
}
