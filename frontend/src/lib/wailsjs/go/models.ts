export namespace database {
	
	export class Opportunity {
	    noticeID?: string;
	    title?: string;
	    solicitationNumber: string;
	    fullParentPathName?: string;
	    fullParentPathCode?: string;
	    postedDate?: string;
	    type?: string;
	    baseType?: string;
	    archiveType?: string;
	    archiveDate?: string;
	    typeOfSetAsideDescription?: string;
	    typeOfSetAside?: string;
	    responseDeadLine?: string;
	    naicsCode?: string;
	    classificationCode?: string;
	    active?: string;
	    description?: string;
	    organizationType?: string;
	    uiLink?: string;
	
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
	        this.classificationCode = source["classificationCode"];
	        this.active = source["active"];
	        this.description = source["description"];
	        this.organizationType = source["organizationType"];
	        this.uiLink = source["uiLink"];
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

}

