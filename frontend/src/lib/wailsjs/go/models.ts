export namespace database {
	
	export class Link {
	    rel: string;
	    href: string;
	
	    static createFrom(source: any = {}) {
	        return new Link(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.rel = source["rel"];
	        this.href = source["href"];
	    }
	}
	export class PointOfContact {
	    fax: string;
	    type: string;
	    email: string;
	    phone: string;
	    title: string;
	    fullName: string;
	
	    static createFrom(source: any = {}) {
	        return new PointOfContact(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fax = source["fax"];
	        this.type = source["type"];
	        this.email = source["email"];
	        this.phone = source["phone"];
	        this.title = source["title"];
	        this.fullName = source["fullName"];
	    }
	}
	export class Opportunity {
	    noticeID: string;
	    title: string;
	    solicitationNumber: string;
	    fullParentPathName: string;
	    fullParentPathCode: string;
	    postedDate: string;
	    type: string;
	    baseType: string;
	    archiveType: string;
	    archiveDate: string;
	    typeOfSetAsideDescription: string;
	    typeOfSetAside: string;
	    responseDeadLine: string;
	    naicsCode: string;
	    naicsCodes: string;
	    classificationCode: string;
	    active: string;
	    pointOfContact: PointOfContact[];
	    description: string;
	    organizationType: string;
	    officeAddress: string;
	    additionalInfoLink: string;
	    uiLink: string;
	    links: Link[];
	    resourceLinks: string;
	
	    static createFrom(source: any = {}) {
	        return new Opportunity(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.noticeID = source["noticeID"];
	        this.title = source["title"];
	        this.solicitationNumber = source["solicitationNumber"];
	        this.fullParentPathName = source["fullParentPathName"];
	        this.fullParentPathCode = source["fullParentPathCode"];
	        this.postedDate = source["postedDate"];
	        this.type = source["type"];
	        this.baseType = source["baseType"];
	        this.archiveType = source["archiveType"];
	        this.archiveDate = source["archiveDate"];
	        this.typeOfSetAsideDescription = source["typeOfSetAsideDescription"];
	        this.typeOfSetAside = source["typeOfSetAside"];
	        this.responseDeadLine = source["responseDeadLine"];
	        this.naicsCode = source["naicsCode"];
	        this.naicsCodes = source["naicsCodes"];
	        this.classificationCode = source["classificationCode"];
	        this.active = source["active"];
	        this.pointOfContact = this.convertValues(source["pointOfContact"], PointOfContact);
	        this.description = source["description"];
	        this.organizationType = source["organizationType"];
	        this.officeAddress = source["officeAddress"];
	        this.additionalInfoLink = source["additionalInfoLink"];
	        this.uiLink = source["uiLink"];
	        this.links = this.convertValues(source["links"], Link);
	        this.resourceLinks = source["resourceLinks"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	
	export class Setting {
	    key: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new Setting(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.value = source["value"];
	    }
	}

}

export namespace main {
	
	export class LoginResponse {
	    message: string;
	    result: boolean;
	
	    static createFrom(source: any = {}) {
	        return new LoginResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.result = source["result"];
	    }
	}
	export class Response {
	    message: string;
	    success: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.success = source["success"];
	    }
	}
	export class SettingRequest {
	    key: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new SettingRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.value = source["value"];
	    }
	}
	export class SettingResponse {
	    message: string;
	    result: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SettingResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.result = source["result"];
	    }
	}

}

export namespace opportunity {
	
	export class OpportunityFilter {
	    fromDate: string;
	    toDate: string;
	    type: string[];
	    naicsCode: string[];
	    page: number;
	    perPage: number;
	
	    static createFrom(source: any = {}) {
	        return new OpportunityFilter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fromDate = source["fromDate"];
	        this.toDate = source["toDate"];
	        this.type = source["type"];
	        this.naicsCode = source["naicsCode"];
	        this.page = source["page"];
	        this.perPage = source["perPage"];
	    }
	}
	export class PaginatedResult {
	    total: number;
	    data: database.Opportunity[];
	
	    static createFrom(source: any = {}) {
	        return new PaginatedResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.total = source["total"];
	        this.data = this.convertValues(source["data"], database.Opportunity);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

export namespace setting {
	
	export class SystemState {
	    lastPull: string;
	
	    static createFrom(source: any = {}) {
	        return new SystemState(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.lastPull = source["lastPull"];
	    }
	}

}

