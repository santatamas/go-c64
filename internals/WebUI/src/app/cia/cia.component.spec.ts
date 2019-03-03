import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CiaComponent } from './cia.component';

describe('CiaComponent', () => {
  let component: CiaComponent;
  let fixture: ComponentFixture<CiaComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CiaComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CiaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
