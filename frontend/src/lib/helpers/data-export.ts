import get from "lodash/get";
import map from "lodash/map";
import forEach from "lodash/forEach";
import type {Header} from "$lib/components/dataTable/datatable";

export function downloadCsvFromArray(data: any[], headers: Header[], filename: string = ''): void {
    let csvData: string = ''
    const rowHeaders: string = map(headers, (o: Header) => o.field).join(',')
    csvData += `${rowHeaders}\n`
    forEach(data, o => {
        const row: any = []
        forEach(headers, (header: Header): void => {
            const value = get(o, header.field) || ''
            row.push(`"${value}"`)
        })
        csvData += `${row.join(',')}\n`
    })

    const blob: Blob = new Blob([csvData], {type: 'text/csv;charset=utf-8,'})
    window.open(  URL.createObjectURL(blob))

    // const hiddenElement: HTMLAnchorElement = createElement('a');
    // hiddenElement.href = URL.createObjectURL(blob);
    // hiddenElement.target = '_blank';
    //
    // //provide the name for the CSV file to be downloaded
    // hiddenElement.download = filename ? `${filename}.csv` : `download.csv`;
    // hiddenElement.click();
}

