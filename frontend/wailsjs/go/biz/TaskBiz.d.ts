// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {entity} from '../models';
import {model} from '../models';
import {biz} from '../models';

export function Add(arg1:entity.Task):Promise<model.HttpResult>;

export function Delete(arg1:string):Promise<model.HttpResult>;

export function Edit(arg1:entity.Task):Promise<model.HttpResult>;

export function ExecStep(arg1:entity.Step):Promise<model.HttpResult>;

export function Get(arg1:string):Promise<model.HttpResult>;

export function GetList():Promise<model.HttpResult>;

export function GetSystemList():Promise<model.HttpResult>;

export function InitBrowser():Promise<void>;

export function Page(arg1:biz.TaskQuery):Promise<model.HttpResult>;
