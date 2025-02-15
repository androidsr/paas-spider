export namespace biz {
	
	export class PlanQuery {
	    page?: model.PageInfo;
	    name: string;
	    execRemark: string;
	
	    static createFrom(source: any = {}) {
	        return new PlanQuery(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = this.convertValues(source["page"], model.PageInfo);
	        this.name = source["name"];
	        this.execRemark = source["execRemark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RecordQuery {
	    page?: model.PageInfo;
	    taskId: string;
	    name: string;
	    system: string;
	
	    static createFrom(source: any = {}) {
	        return new RecordQuery(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = this.convertValues(source["page"], model.PageInfo);
	        this.taskId = source["taskId"];
	        this.name = source["name"];
	        this.system = source["system"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class TaskQuery {
	    page?: model.PageInfo;
	    name: string;
	    system: string;
	
	    static createFrom(source: any = {}) {
	        return new TaskQuery(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = this.convertValues(source["page"], model.PageInfo);
	        this.name = source["name"];
	        this.system = source["system"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace entity {
	
	export class Config {
	    id: string;
	    legalStatement: string;
	    content: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.legalStatement = source["legalStatement"];
	        this.content = source["content"];
	    }
	}
	export class Plan {
	    id: string;
	    name: string;
	    execCron: string;
	    execRemark: string;
	    content: string;
	    status: string;
	
	    static createFrom(source: any = {}) {
	        return new Plan(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.execCron = source["execCron"];
	        this.execRemark = source["execRemark"];
	        this.content = source["content"];
	        this.status = source["status"];
	    }
	}
	export class Record {
	    id: string;
	    taskId: string;
	    selector: string;
	    inputValue: string;
	    outputValue: string;
	    execTime: string;
	    sourceType: string;
	
	    static createFrom(source: any = {}) {
	        return new Record(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.taskId = source["taskId"];
	        this.selector = source["selector"];
	        this.inputValue = source["inputValue"];
	        this.outputValue = source["outputValue"];
	        this.execTime = source["execTime"];
	        this.sourceType = source["sourceType"];
	    }
	}
	export class Step {
	    id: string;
	    taskId: string;
	    name: string;
	    eventType: string;
	    selector: string;
	    inputValue: string;
	    sleepTime: number;
	
	    static createFrom(source: any = {}) {
	        return new Step(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.taskId = source["taskId"];
	        this.name = source["name"];
	        this.eventType = source["eventType"];
	        this.selector = source["selector"];
	        this.inputValue = source["inputValue"];
	        this.sleepTime = source["sleepTime"];
	    }
	}
	export class Task {
	    id: string;
	    name: string;
	    system: string;
	    publicId: string;
	    execTime: string;
	    content: string;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.system = source["system"];
	        this.publicId = source["publicId"];
	        this.execTime = source["execTime"];
	        this.content = source["content"];
	    }
	}

}

export namespace model {
	
	export class HttpResult {
	    code: number;
	    msg: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new HttpResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.msg = source["msg"];
	        this.data = source["data"];
	    }
	}
	export class OrderItem {
	    column: string;
	    asc: boolean;
	
	    static createFrom(source: any = {}) {
	        return new OrderItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.column = source["column"];
	        this.asc = source["asc"];
	    }
	}
	export class PageInfo {
	    current: number;
	    size: number;
	    orders: OrderItem[];
	
	    static createFrom(source: any = {}) {
	        return new PageInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.current = source["current"];
	        this.size = source["size"];
	        this.orders = this.convertValues(source["orders"], OrderItem);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

