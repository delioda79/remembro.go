package main

import "strconv"

/*
Protocol handles the format
*/
type Protocol struct {
	Commands     []string
	length       int
	lengthString string
}

func (p *Protocol) check(text string) error {
	if p.length == 0 {
		pos := 0
		for {
			if string(text[pos]) == "{" {
				break
			}
			_, err := strconv.Atoi(string(text[pos]))
			if err != nil {
				p.length = 0
				p.lengthString = ""
				return err
			}
			p.lengthString += string(text[pos])
			if pos == len(text)-1 {
				return nil
			}
			pos++
		}

		length, _ := strconv.Atoi(p.lengthString)

		p.length = length
		p.lengthString = ""
		newCommands := make([]string, len(p.Commands)+1)
		copy(newCommands, p.Commands)
		p.Commands = newCommands
		return p.check(text[pos:])

	}

	if p.length >= len(text) {
		p.Commands[len(p.Commands)-1] += text
		p.length -= len(text)
		return nil
	}
	first := text[:p.length]
	last := text[p.length:len(text)]
	p.Commands[len(p.Commands)-1] += first
	p.length = 0
	return p.check(last)
}
