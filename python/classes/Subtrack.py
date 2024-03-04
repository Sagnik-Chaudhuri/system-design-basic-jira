class Subtrack():

    allowed_status = ['open', 'in_progress', 'completed']

    def __init__(self, subtrack_title, status, parent_task_title) -> None:
        self.subtrack_title = subtrack_title
        self.status = status if status in self.allowed_status else None
        self.parent_task_title = parent_task_title

    def display_details(self):
        print('\nsubtrack title: ', self.subtrack_title, '\nsubtrack status',
              self.status, '\nparent task title: ', self.parent_task_title)

    def change_status(self, new_status):
        if new_status not in self.allowed_status:
            print('invalid subtrack status')
            return
        if self.allowed_status.index(new_status) < self.allowed_status.index(self.status):
            print('new status cannot be behind prevous status')
            return

        self.status = new_status
        return
