import type * as monacoEditor from 'monaco-editor';

export interface EditorInput {
    text: string;
    language?: string;
}

export interface MonacoModelFactory {
    editor: {
        createModel(value: string, language?: string): monacoEditor.editor.ITextModel;
    };
}

export function isEqualInput(first?: EditorInput, second?: EditorInput) {
    return first && second && first.text === second.text && (first.language || '') === (second.language || '');
}

// Replace a document's contents with text the user must not be able to get back. setValue discards
// the model's edit history, so undo cannot resurrect a stale refresh or an abandoned edit and save it
// over the cluster. An edit operation would keep both on the undo stack. Callers are expected to
// restore the view state afterwards, since setValue resets it.
//
// This deliberately does not go through editor.executeEdits, which does nothing while the editor is
// readOnly, and the manifest view is readOnly whenever it is not being edited.
export function replaceModelText(model: monacoEditor.editor.ITextModel, text: string): void {
    model.setValue(text);
}

// Update an existing Monaco editor's document without treating a live-text refresh as a new file.
// Only the incoming props are compared, never the buffer, so text the user is still typing survives.
export function applyEditorInput(monaco: MonacoModelFactory, editor: monacoEditor.editor.IStandaloneCodeEditor, prev: EditorInput | undefined, next: EditorInput): void {
    if (isEqualInput(prev, next)) {
        return;
    }

    const viewState = editor.saveViewState();
    const model = editor.getModel();
    const languageChanged = (prev?.language || '') !== (next.language || '');

    if (model && !languageChanged) {
        replaceModelText(model, next.text);
    } else {
        const newModel = monaco.editor.createModel(next.text, next.language);
        editor.setModel(newModel);
        model?.dispose();
    }

    if (viewState) {
        editor.restoreViewState(viewState);
    }
}
