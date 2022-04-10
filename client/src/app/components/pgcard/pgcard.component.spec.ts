import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PgcardComponent } from './pgcard.component';

describe('PgcardComponent', () => {
  let component: PgcardComponent;
  let fixture: ComponentFixture<PgcardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PgcardComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PgcardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
