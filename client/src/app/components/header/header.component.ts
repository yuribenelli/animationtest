import { Component, OnInit } from '@angular/core';
import {MatInputModule} from '@angular/material/input';
import { Character } from 'src/app/model/character.model';
import {MatFormFieldModule} from '@angular/material/form-field';
import { HtmlParser } from '@angular/compiler';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {

  newCharacter? : Character;


  constructor() { }

  ngOnInit(): void {


  }
  onClick(){




  }




}
