import { Component, OnInit } from '@angular/core';
import { TelemetryService } from '../services/telemetry.service';
import { TablehelperService } from '../services/tablehelper.service';
import { Telemetry } from '../models/telemetry.model';

@Component({
  selector: 'app-cia',
  templateUrl: './cia.component.html',
  styleUrls: ['./cia.component.scss']
})
export class CiaComponent implements OnInit {

  public displayedColumns: string[] = ['name', 'value'];
  public dataSource: any;

  constructor(private telemetryService: TelemetryService, private helper: TablehelperService) {
    telemetryService.getTelemetry().subscribe((t: string) => {
      const telemetry: Telemetry = JSON.parse(t);
      if (telemetry.Command === 'GetCIAState') {
        this.dataSource = helper.convertToTableRowsWithBlackList(JSON.parse(atob(telemetry.Payload)), ['Keyboard_matrix']);
      }
    });
   }

  ngOnInit() {
  }

}
