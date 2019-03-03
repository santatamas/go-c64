import { Injectable } from '@angular/core';
import { TableRow } from '../models/tablerow.model';
import { convertActionBinding } from '@angular/compiler/src/compiler_util/expression_converter';

@Injectable({
  providedIn: 'root'
})
export class TablehelperService {

  constructor() { }

  public convertToTableRows(input: any): TableRow[] {
    return this.convertToTableRowsWithBlackList(input, Array());
  }

  public convertToTableRowsWithBlackList(input: any, blackList: string[]): TableRow[] {

    let result = [];

     // Step 1. Get all the object keys.
     const properties = Object.keys(input);
     // Step 3. Iterate throw all keys.
    for (const prop of properties) {
      if (!blackList.includes(prop)) {
        const row = new TableRow();
        row.Name = prop.toString();

        let value = input[prop];

        if (typeof (value) === 'number') {
          value = value.toString(16);
        }

        row.Value = value;
        result.push(row);
      }
    }
    return result;
  }
}
