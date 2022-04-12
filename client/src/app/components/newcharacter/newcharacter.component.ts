import { Component, Input, OnInit } from '@angular/core';
import { Character } from 'src/app/model/character.model';
import { NgForm } from '@angular/forms';



@Component({
  selector: 'app-newcharacter',
  templateUrl: './newcharacter.component.html',
  styleUrls: ['./newcharacter.component.scss']
})
export class NewcharacterComponent implements OnInit {

  newCharacter = {} as Character;

  constructor() { }

  ngOnInit(): void {

  }

  onAddCharacter(form: NgForm ){

    const value = form.value

    console.log(this.newCharacter.Name)
  }

}
