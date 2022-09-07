package views

import (
	"errors"
	"fmt"
	"time"

	"filec2/config"

	"filec2/db"

	"filec2/Handlers/sessions"
)

var (
	//ErrUnauthorised is returned when the user enters incorrect login details
	ErrUnauthorised = errors.New("user unauthorised")
)

//LoginPage will render the login page and authenticate the user
func LoginPage(session *sessions.Session) error {

	for attempts := 0; attempts <= config.MaxLoginTries; attempts++ {

		if err := session.Clear(); err != nil {
			return err
		}    
    session.Println("[19;61H\x1b[38;5;242m     PASSWORD     ")
    session.Println("[20;61H\x1b[38;5;242m╔════════════════╗")
    session.Println("[21;61H\x1b[38;5;242m║                \x1b[38;5;242m║") 
    session.Println("[22;61H\x1b[38;5;242m╚════════════════╝")
    session.Println("[19;0H\x1b[38;5;242m      USERNAME    ")
    session.Println("[20;0H\x1b[38;5;242m╔════════════════╗")
    session.Println("[21;0H\x1b[38;5;242m║                \x1b[38;5;242m║")  
    session.Println("[22;0H\x1b[38;5;242m╚════════════════╝")
    
    username, err := session.ReadCustom("[21;2H", false, 20)
    if err != nil {
        return err
    }

    password, err := session.ReadCustom("[21;62H", true, 20)
    if err != nil {
        return err
    }

 

  if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}
			if _, err := session.Println("\x1b[38;5;127;8;8m LOADING Sky NET!"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}

			if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}
			if _, err := session.Println("\x1b[38;5;127;8;8m Loading 10%"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}


			if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}
			if _, err := session.Println("\x1b[38;5;127;8;8m Loading 20%"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}


         if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}
			if _, err := session.Println("\x1b[38;5;127;8;8m Loading 40%"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}

         if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}
			if _, err := session.Println("\x1b[38;5;127;8;8m Loading 47%"); err != nil {
				return err
			}
         
			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}
         
         if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}
			if _, err := session.Println("\x1b[38;5;127;8;8m Loading 50%"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}

			if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}
			if _, err := session.Println("\x1b[38;5;127;8;8m Loading 60%"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}


			if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}
			if _, err := session.Println("\x1b[38;5;127;8;8m Loading 70%"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}


			if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}
			if _, err := session.Println("\x1b[38;5;127;8;8m Loading 77%"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}


			if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}
			if _, err := session.Println("\x1b[38;5;127;8;8m Loading 82%"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}


         if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}
			if _, err := session.Println("\x1b[38;5;127;8;8m Loading 83%"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}

         if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}
			if _, err := session.Println("\x1b[38;5;127;8;8m Loading 88%"); err != nil {
				return err
			}
			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}
         
         if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}
			if _, err := session.Println("\x1b[38;5;127;8;8m Loading 90%"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}

			if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}
			if _, err := session.Println("\x1b[38;5;127;8;8m Loading 96%"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}


			if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}
			if _, err := session.Println("\x1b[38;5;127;8;8m Loading 100%"); err != nil {
				return err
			}
         if _, err := session.Println("\x1b[38;5;127;8;8m "); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
         
         
         


   

		user, err := db.GetUser(username)
		if err != nil {
			return err
		}

		if user == nil {
      session.Println("\x1b[0m\x1b[38;2;14;161;129m╔════════════════════╗\r\n")
			session.Println("\x1b[0m\x1b[38;2;14;161;129m║\x1b[38;5;127;8;8mInvalid Credentials.║\r\n")
			session.Println("\x1b[0m\x1b[38;2;14;161;129m╚════════════════════╝\r\n")
		} else {
			if db.HashPassword(password) == user.Password {
				session.User = user

				if user.Banned {
					return Userbanned(session)
				}

				return nil
			}
		}
		if err := session.Clear(); err != nil {
			return err
		}
		  session.Println("\x1b[0m\x1b[38;2;14;161;129m╔════════════════════╗\r\n")
			session.Println("\x1b[0m\x1b[38;2;14;161;129m║\x1b[38;5;127;8;8mInvalid Credentials.║\r\n")
			session.Println("\x1b[0m\x1b[38;2;14;161;129m╚════════════════════╝\r\n")
		if attempts == config.MaxLoginTries-1 {
			break
		}

		cooldown := 3
		for i := 0; i <= cooldown; i++ {
			session.Write([]byte(fmt.Sprintf("\x1b[38;5;127;8;8m ", cooldown-i)))
			time.Sleep(time.Second)
		}

	}

	return LoginTooManyAttempts(session)
}