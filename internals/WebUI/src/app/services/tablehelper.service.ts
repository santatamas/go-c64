import { Injectable } from '@angular/core';
import { TableRow } from '../models/tablerow.model';

@Injectable({ providedIn: 'root' })
export class TablehelperService {
  convertToTableRows(input: any): TableRow[] {
    return this.convertToTableRowsWithBlackList(input, []);
  }

  convertToTableRowsWithBlackList(input: any, blackList: string[]): TableRow[] {
    const result: TableRow[] = [];
    const properties = Object.keys(input);
    for (const prop of properties) {
      if (!blackList.includes(prop)) {
        let value = input[prop];
        if (typeof value === 'number') {
          value = value.toString(16);
        }
        result.push({ Name: prop, Value: value });
      }
    }
    return result;
  }
}
