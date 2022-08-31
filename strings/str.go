package main

import "strings"

func removeOutermostParenthesis(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))
	var depth int

	for _, r := range s {

		if r == '(' && depth > 0 || r == ')' && depth > 1 {
			sb.WriteRune(r)
		}

		if r == '(' {
			depth++
		} else {
			depth--
		}
	}
	return sb.String()
}

func reverseWords(s string) string {
    str := strings.TrimSpace(s)
    word := strings.Fields(str)
    
    for i,j:=0, len(word)-1; i<j; i,j= i+1,j-1 {
        word[i], word[j] = word[j], word[i]
    }
    return strings.Join(word, " ")
}

func largestOddNumber(num string) string {
    
    var i = len(num)-1
    
    for i >=0 && num[i]%2 == 0 {
        i--
    }
    return num[:i+1]
}

func longestCommonPrefix(strs []string) string {
    var ans strings.Builder
    
    for i := 0; i < len(strs[0]); i++ {
        cur := strs[0][i]
        for j := 1; j < len(strs); j++ {
            if i == len(strs[j]) || strs[j][i] != cur {
                return ans.String()
            }
        }
        ans.WriteString(string(cur))
    }
    
    return ans.String()
}

func isIsomorphic(s string, t string) bool {
    
    if len(s) != len(t) {
        return false
    }
		smap:=make(map[byte]byte)  //*using byte type to avoid unneccesary byte to string conversion*
		tmap:=make(map[byte]byte)
		for i:=0;i<len(s);i++{
			if _,ok:=smap[s[i]]; !ok{
				smap[s[i]]=t[i]  //*If we dont find the element in map, we add them mapping byte code of character from 's' to that of 't' as key , value respectively*
			}
			if _,ok:=tmap[t[i]]; !ok{
				tmap[t[i]]=s[i] //* Similarily, mapping byte code of character from 't' to that of 's' as key , value respectively*
			}
			if smap[s[i]]!=t[i] || tmap[t[i]]!=s[i]{
				return false // If at any point we come across these characters are not mapped one-to-one i.e one character from s is mapped to two different characters of t or viceversa, we return False
			}
		}
		return true 
    // If we dont come across the above condition , that means our input has one-to-one mapped character //set from 's' and 't' strings, hence we return True*
}

func rotateString(s string, goal string) bool {
    
    if len(s) != len(goal) {
        return false
    }
    
    return strings.Contains(goal+goal, s)
}

func isAnagram(s string, t string) bool {
    
    var lenS = len(s)
    var lenT = len(t)
    
    if lenS != lenT {
        return false
    }
    
    //map to store each character and count number of it
    var sm = make(map[string]int) 
    
    for _,v:=range s {
        sm[string(v)]++
    }
    
    for _,v:=range t {
        sm[string(v)]--
    }
    
    for _,v:=range s {
        if sm[string(v)] != 0 {
            return false
        }
    }
    
    return true
}