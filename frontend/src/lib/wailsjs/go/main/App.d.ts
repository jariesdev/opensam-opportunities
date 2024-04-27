// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {database} from '../models';
import {setting} from '../models';
import {main} from '../models';
import {opportunity} from '../models';

export function GetAllSettings():Promise<Array<database.Setting>>;

export function GetCwd():Promise<string>;

export function GetOpportunityNaicsCodes():Promise<Array<string>>;

export function GetOpportunityTypes():Promise<Array<string>>;

export function GetSystemStatus():Promise<setting.SystemState>;

export function Greet(arg1:string):Promise<string>;

export function Login(arg1:string,arg2:string):Promise<main.LoginResponse>;

export function PullLatest():Promise<void>;

export function SearchOpportunities(arg1:string,arg2:opportunity.OpportunityFilter,arg3:boolean):Promise<opportunity.PaginatedResult>;

export function UpdateSettings(arg1:Array<main.SettingRequest>):Promise<main.SettingResponse>;
