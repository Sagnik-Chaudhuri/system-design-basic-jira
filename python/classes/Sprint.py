class Sprint():
    allowed_status = ['open', 'in_progress', 'completed']

    def __init__(self, name, start_date, end_date, status, tasks=[]) -> None:
        self.name = name
        self.start_date = start_date
        self.end_date = end_date
        self.status = status if status in self.allowed_status else None
        self.tasks = tasks

    def add_task_to_sprint(self, task):
        if self.status != 'completed':
            self.tasks.append(task)
        else:
            print('cant add to sprint, sprint ended already')

    def remove_task_from_sprint(self, task_title):
        if (len(self.tasks) == 0):
            print('\ncannot remove task from empty sprint')
            return
        for index, task in enumerate(self.tasks):
            if task.title == task_title:
                del self.tasks[index]
                print('\ntask removed from sprint')

    def change_status(self, new_status):
        if new_status not in self.allowed_status:
            print('invalid bug status')
            return
        if self.allowed_status.index(new_status) < self.allowed_status.index(self.status):
            print('new status cannot be behind prevous status')
            return

        print('\nchanging status of sprint from',
              self.status, 'to ', new_status)
        self.status = new_status
        return

    def start_sprint(self):
        self.change_status('open')

    def end_sprint(self):
        self.change_status('completed')

    def display_sprint_details(self):
        print('\ndisplaying sprint details')
        print('\nsprint name', self.name, '\nsprint start_date: ', self.start_date, '\nend_date: ', self.end_date,
              '\nstatus: ', self.status)
        if (len(self.tasks) == 0):
            print('\nsprint has no tasks')
            return
        for task in self.tasks:
            task.display_details()
