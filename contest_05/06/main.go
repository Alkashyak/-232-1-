
    type Cat struct {
        alive bool 
    }
    type Box struct{
    }
    func (b Box) open() Cat {
    gj := false
    a := rand.Intn(4)
    if a >= 2 {
        gj = true
    }
    return cca(gj)
	}
    func (n2 Cat) is_alive()  bool {
        return n2.alive
    }
    func cca (b bool) Cat{
        acc := Cat{
            alive: b,
        }
        return acc
    }
    
