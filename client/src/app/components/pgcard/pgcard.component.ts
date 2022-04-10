import { Component, Input } from '@angular/core';
import { Character } from 'src/app/model/character.model';


@Component({
  selector: 'app-pgcard',
  templateUrl: './pgcard.component.html',
  styleUrls: ['./pgcard.component.scss']
})
export class PgcardComponent  {
  @Input() character? : Character;
}
