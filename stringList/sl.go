package stringlist

import "errors"

// StringList: односвязный список, представляющий строку.
type StringList struct {
	Head *Node
	Tail *Node
	Length int
}


// Создает новый StringList из обычной строки Go.
func New(s string) *StringList{
	var result StringList
	for _, c := range s {
		result.Append(c)
	}
	return &result
}

// Возвращает длину строки.
func (s *StringList) Len() int {
	return s.Length
}

// Добавляет символ в конец строки.
func (s *StringList) Append(c rune){
	if s.Length == 0 {
		s.Head = &Node{Symbol: c}
		s.Tail = s.Head
	} else {
		s.Tail.Next = &Node{Symbol: c}
		s.Tail = s.Tail.Next
	}
	s.Length++
}

// Добавляет символ в начало строки.
func (s *StringList) Prepend(c rune){
	if s.Length == 0 {
		s.Head = &Node{Symbol: c}
		s.Tail = s.Head
	} else {
		s.Head = &Node{Symbol: c, Next: s.Head}
	}
	s.Length++
}

// Вставляет символ на определенную позицию в строке.
func (s *StringList) Insert(c rune, index int) error{
	if index > s.Length {
		return errors.New("index out of range")
	}
	if index == 0 {
		s.Prepend(c)
	} else if index == s.Length {
		s.Append(c)
	} else {
		prev := s.Head
		for i := 0; i < index-1; i++ {
			prev = prev.Next
		}
		prev.Next = &Node{Symbol: c, Next: prev.Next}
		s.Length++
	}
	return nil
}

// Удаляет символ из строки по индексу.
func (s *StringList) Remove(index int) error {
	if index >= s.Length {
		return errors.New("index out of range")
	}
	if index == 0 {
		s.Head = s.Head.Next
	} else {
		prev := s.Head
		for i := 0; i < index-1; i++ {
			prev = prev.Next
		}
		prev.Next = prev.Next.Next
	}
	s.Length--
	return nil
}

// Возвращает строковое представление списка.
func (s *StringList) String() string {
	result := ""
	current := s.Head
	for current != nil {
		result += string(current.Symbol)
		current = current.Next
	}
	return result
}

// Получает символ по индексу.
func (s *StringList) At(index int) (rune, error){
	if index >= s.Length {
		return 0, errors.New("index out of range")
	}
	current := s.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Symbol, nil
}

// Конкатенирует две строки.
func (s *StringList) Concat(other *StringList) *StringList{
	result := New("")
	current := s.Head
	for current != nil {
		result.Append(current.Symbol)
		current = current.Next
	}
	current = other.Head
	for current != nil {
		result.Append(current.Symbol)
		current = current.Next
	}
	return result
}

// Проверяет, равны ли две строки.
func (s *StringList) Equals(other *StringList) bool {
	if s.Length != other.Length {
		return false
	}
	current1 := s.Head
	current2 := other.Head
	for current1 != nil {
		if current1.Symbol != current2.Symbol {
			return false
		}
		current1 = current1.Next
		current2 = current2.Next
	}
	return true
}

// Возвращает первый индекс указанного символа в строке.
func (s *StringList) IndexOf(c rune) int {
	current := s.Head
	for i := 0; i < s.Length; i++ {
		if current.Symbol == c {
			return i
		}
		current = current.Next
	}
	return -1
}

// Возвращает подстроку из строки.
func (s *StringList) Substring(start, end int) (*StringList, error) {
	if start >= s.Length || end >= s.Length {
		return nil, errors.New("index out of range")
	}
	result := New("")
	current := s.Head
	for i := 0; i < start; i++ {
		current = current.Next
	}
	for i := start; i < end; i++ {
		result.Append(current.Symbol)
		current = current.Next
	}
	return result, nil
}

// Заменяет символы в строке.
func (s *StringList) Replace(old, new rune, n int) *StringList {
	result := New("")
	current := s.Head
	for i := 0; i < s.Length; i++ {
		if current.Symbol == old && n > 0 {
			result.Append(new)
			n--
		} else {
			result.Append(current.Symbol)
		}
		current = current.Next
	}
	return result
}
