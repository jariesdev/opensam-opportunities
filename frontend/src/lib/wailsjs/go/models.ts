export namespace opportunities {
	
	export class  {
	    rel: string;
	    href: string;
	
	    static createFrom(source: any = {}) {
	        return new (source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.rel = source["rel"];
	        this.href = source["href"];
	    }
	}
	export class  {
	    fax: any;
	    type: string;
	    email: string;
	    phone: string;
	    title: any;
	    fullName: string;
	
	    static createFrom(source: any = {}) {
	        return new (source);
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
	export class OpportunityData {
	    noticeId: string;
	    title: string;
	    solicitationNumber: string;
	    fullParentPathName: string;
	    fullParentPathCode: string;
	    postedDate: string;
	    type: string;
	    baseType: string;
	    archiveType: string;
	    archiveDate: string;
	    typeOfSetAsideDescription: any;
	    typeOfSetAside: any;
	    responseDeadLine: string;
	    naicsCode: string;
	    naicsCodes: string[];
	    classificationCode: string;
	    active: string;
	    award: any;
	    pointOfContact: [];
	    description: string;
	    organizationType: string;
	    // Go type: struct { Zipcode string "json:\"zipcode\""; City string "json:\"city\""; CountryCode string "json:\"countryCode\""; State string "json:\"state\"" }
	    officeAddress: any;
	    // Go type: struct { City struct { Code string "json:\"code\""; Name string "json:\"name\"" } "json:\"city\""; State struct { Code string "json:\"code\""; Name string "json:\"name\"" } "json:\"state\""; Country struct { Code string "json:\"code\""; Name string "json:\"name\"" } "json:\"country\"" }
	    placeOfPerformance: any;
	    additionalInfoLink: any;
	    uiLink: string;
	    links: [];
	    resourceLinks: string[];
	
	    static createFrom(source: any = {}) {
	        return new OpportunityData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.noticeId = source["noticeId"];
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
	        this.award = source["award"];
	        this.pointOfContact = this.convertValues(source["pointOfContact"], );
	        this.description = source["description"];
	        this.organizationType = source["organizationType"];
	        this.officeAddress = this.convertValues(source["officeAddress"], Object);
	        this.placeOfPerformance = this.convertValues(source["placeOfPerformance"], Object);
	        this.additionalInfoLink = source["additionalInfoLink"];
	        this.uiLink = source["uiLink"];
	        this.links = this.convertValues(source["links"], );
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
	export class ResultLink {
	    rel: string;
	    href: string;
	
	    static createFrom(source: any = {}) {
	        return new ResultLink(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.rel = source["rel"];
	        this.href = source["href"];
	    }
	}
	export class SearchResult {
	    totalRecords: number;
	    limit: number;
	    offset: number;
	    opportunitiesData: OpportunityData[];
	    links: ResultLink[];
	
	    static createFrom(source: any = {}) {
	        return new SearchResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.totalRecords = source["totalRecords"];
	        this.limit = source["limit"];
	        this.offset = source["offset"];
	        this.opportunitiesData = this.convertValues(source["opportunitiesData"], OpportunityData);
	        this.links = this.convertValues(source["links"], ResultLink);
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

