import { Component } from '@angular/core';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { ControlsComponent } from './controls/controls.component';
import { CpuComponent } from './cpu/cpu.component';
import { CiaComponent } from './cia/cia.component';
import { EmulatorComponent } from './emulator/emulator.component';
import { MemoryComponent } from './memory/memory.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    MatToolbarModule,
    MatIconModule,
    ControlsComponent,
    CpuComponent,
    CiaComponent,
    EmulatorComponent,
    MemoryComponent
  ],
  template: `
    <mat-toolbar color="primary">
      <mat-toolbar-row>
        <app-controls></app-controls>
        <span class="spacer"></span>
        <mat-icon>memory</mat-icon>
      </mat-toolbar-row>
    </mat-toolbar>
    <div class="content">
      <app-emulator></app-emulator>
      <app-cpu></app-cpu>
      <app-cia></app-cia>
      <app-memory></app-memory>
    </div>
  `,
  styles: [`
    .spacer { flex: 1 1 auto; }
    .content { padding: 20px; display: flex; flex-direction: column; gap: 16px; }
  `]
})
export class AppComponent {}
