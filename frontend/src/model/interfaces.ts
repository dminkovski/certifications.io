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
    notes: string;
    image: string;
    skills: string[];
    link: string;
    courses:  ICourse[];
}