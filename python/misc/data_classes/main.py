from dataclasses import dataclass, field

@dataclass(order=True, frozen=True)
class Person:
    sort_index: int = field(init=False, repr=False)
    name: str
    job: str
    age: int
    strength: int = 100
    
    def __post_init__(self):
        # self.sort_index = self.age
        object.__setattr__(self, 'sort_index', self.strength)
        
    def __repr__(self):
        return f'{self.name}, {self.job} ({self.age})'
    
p1 = Person("Geralt", "Witcher", 30, 200)
p2 = Person("Yennefer", "Sorceress", 25)
p3 = Person("Yennefer", "Sorceress", 25)

print(id(p2), id(p3)) # 4315216384 4315216816
print(p2 == p3) # True
print(p1 > p2) # True
print(p1) # Geralt, Witcher (30)