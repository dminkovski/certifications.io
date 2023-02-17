export interface IProvider {
    name: string;
    link: string;
    id: number;
}

export interface ICourse {
    name: string;
    provider: IProvider;
    link: string;
}

export interface ICertification {
    id:number;
    name: string;
    image: string;
    description: string;
    skills: string[];
    link: string;
    courses:  ICourse[];
    minMonths: number;
    maxMonths: number;
    shortName: string;
    notes: string;
    updated: string | Date;
    created: string | Date;
    price: number;
}