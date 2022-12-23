export interface IProvider {
    Name: string;
    Link: string;
    Id: number;
}

export interface ICourse {
    Name: string;
    Provider: IProvider;
    Link: string;
}

export interface ICertification {
    Id:number;
    Name: string;
    Notes: string;
    Image: string;
    Skills: string[];
    Link: string;
    Courses:  ICourse[];
}