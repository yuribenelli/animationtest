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

    this.newCharacter.Name = value.Name
    this.newCharacter.Race = value.race
    this.newCharacter.Class = value.class
    this.newCharacter.BaseDamage = value.base_damage
    this.newCharacter.Hp = value.hp
    this.newCharacter.Background = value.background
    console.log(this.newCharacter.Name)
    console.log(this.newCharacter.Race)
    console.log(this.newCharacter.Class)
    console.log(this.newCharacter.BaseDamage)
    console.log(this.newCharacter.Hp)
    console.log(this.newCharacter.Background)
  }

}
