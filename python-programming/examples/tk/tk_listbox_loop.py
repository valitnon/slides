import tkinter as tk

app = tk.Tk()
app.title('List box')


def clicked():
   print("clicked")
   selected = box.curselection()  # returns a tuple
   print(box.get(selected[0]))
   exit()

box = tk.Listbox(app)
values = ['Red', 'Green', 'Blue', 'Purple']
for val in values:
   box.insert(tk.END, val)
box.pack()

button = tk.Button(app, text='Close', width=25, command=clicked)
button.pack()

app.mainloop()

