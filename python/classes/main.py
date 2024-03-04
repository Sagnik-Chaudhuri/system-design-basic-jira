from classes.Sprint import Sprint
from classes.Subtrack import Subtrack
from classes.Story import Story
from classes.Bug import Bug
from classes.Feature import Feature
from classes.Task import Task


all_tasks_grouped_by_user = {}
all_tasks_list = []
sprints_list = []


def add_task(title, assignee, task_type, due_date):
    if task_type == 'feature':
        print('\nthis is a feature task')
        feature_summary = input('feature summary?')
        impact = input('impact?')
        status = input('feature status?')

        feature = Feature(title=title, assignee=assignee, task_type='feature', due_date=due_date,
                          feature_summary=feature_summary, impact=impact, status=status)
        if all_tasks_grouped_by_user.get(assignee) is None:
            all_tasks_grouped_by_user[assignee] = [title]
        else:
            all_tasks_grouped_by_user[assignee].append(title)
        all_tasks_list.append(feature)

    elif task_type == 'bug':
        print('\nthis is a bug task')
        severity = input('\nseverity?')
        status = input('\nstatus?')
        bug = Bug(title=title, assignee=assignee, task_type=task_type,
                  due_date=due_date, severity=severity, status=status)
        if all_tasks_grouped_by_user.get(assignee) is None:
            all_tasks_grouped_by_user[assignee] = [title]
        else:
            all_tasks_grouped_by_user[assignee].append(title)
        all_tasks_list.append(bug)

    elif task_type == 'story':
        print('\nthis is story')
        story_summary = input('story summary?')
        status = input('status?')

        story = Story(title=title, assignee=assignee, task_type=task_type, due_date=due_date,
                      story_summary=story_summary, status=status)
        if all_tasks_grouped_by_user.get(assignee) is None:
            all_tasks_grouped_by_user[assignee] = [title]
        else:
            all_tasks_grouped_by_user[assignee].append(title)
        all_tasks_list.append(story)


def change_task_status(task_title, new_status):
    for task in all_tasks_list:
        if task.title == task_title:
            task.change_status(new_status)


def change_assignee(task_title, new_assignee):
    old_assignee = None
    for task in all_tasks_list:
        if task.title == task_title:
            old_assignee = task.assignee
            task.change_task_assignee(new_assignee)

    if all_tasks_grouped_by_user.get(new_assignee) is None:
        all_tasks_grouped_by_user[new_assignee] = [task.title]
    else:
        all_tasks_grouped_by_user[new_assignee].append(task.title)

    all_tasks_grouped_by_user[old_assignee].remove(task.title)

    return task


def display_tasks_assigned_to_user(username):
    print('\nusername received: ', username)
    if all_tasks_grouped_by_user.get(username) is None:
        print('\nNo task assigned to current username.')
        return
    else:
        print('displaying tasks for user: ', username)
        for task in all_tasks_list:
            if task.assignee == username and task.status is not None:
                task.display_details()


def add_subtrack_to_existing_story(title):
    for task in all_tasks_list:
        if task.title == title and task.task_type == 'story':
            print('add subtrack?')
            task.add_subtrack_to_current_story()


def change_subtrack_status(task_title, new_status, subtrack_title):
    for task in all_tasks_list:
        if task.title == task_title:
            for subtrack in task.subtracks:
                if subtrack.subtrack_title == subtrack_title:
                    subtrack.change_status(new_status)


def add_sprint():
    print('\ninput sprint details: ')
    name = input('\nsprint name?')
    start_date = input('sart date?')
    end_date = input('\nend_date?')
    status = input('\nstatus?')
    sprint = Sprint(name, start_date, end_date, status)
    sprints_list.append(sprint)


def add_task_to_sprint(task_title, sprint_name):
    task_found = False
    sprint_found = False
    for sprint in sprints_list:
        if sprint.name == sprint_name:
            sprint_found = True
            for task in all_tasks_list:
                if task.title == task_title:
                    task_found = True
                    sprint.add_task_to_sprint(task)
                    return

    if not sprint_found:
        print('\nsprint not found')
    elif not task_found:
        print('\ntask not found')


def remove_task_from_sprint(task_title, sprint_name):
    task_found = False
    for sprint in sprints_list:
        if sprint.name == sprint_name:
            sprint_found = True
            for task in sprint.tasks:
                if task.title == task_title:
                    sprint.remove_task_from_sprint(task_title)
                    task_found = True
                    return

    if not sprint_found:
        print('\nsprint not found')
    elif not task_found:
        print('\ntask not found')


def start_sprint(sprint_name):
    for sprint in sprints_list:
        if sprint.name == sprint_name:
            sprint.start_sprint()
            return
    print('\nsprint not found')


def end_sprint(sprint_name):
    for sprint in sprints_list:
        if sprint.name == sprint_name:
            sprint.end_sprint()
            return
    print('\nsprint not found')


def display_sprint_snapshot(sprint_name):
    for sprint in sprints_list:
        if sprint_name == sprint.name:
            sprint.display_sprint_details()
            return

    print('sprint not found')


def main():

    print('\n\nmain start here\n\n')
    add_task('task_1', 'peter', 'story', 'tom')
    display_tasks_assigned_to_user('peter')
    # print('all_tasks', all_tasks_list)
    # change_task_status('abcd', 'open')
    # # add_subtrack_to_existing_story('abcd')
    # display_tasks_assigned_to_user('peter')
    add_sprint()
    # end_sprint('sprint_1')
    add_task_to_sprint('task_1', 'sprint_1')
    display_sprint_snapshot('sprint_1')
    add_subtrack_to_existing_story('task_1')
    display_sprint_snapshot('sprint_1')
    display_tasks_assigned_to_user('peter')
    # change_assignee('task_1', 'raju')

    # display_tasks_assigned_to_user('raju')

    return


main()
